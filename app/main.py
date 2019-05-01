import os
from flask import Flask, render_template, flash, request, redirect, url_for, session
import couchdb
import json

app = Flask(__name__)

app.config['COUCHDB_USER'] = os.getenv("COUCHDB_USER")
app.config['COUCHDB_PASSWORD'] = os.getenv("COUCHDB_PASSWORD")
app.secret_key = os.getenv("SECRET_KEY")

# Init couchdb
couchconn = couchdb.Server("http://%s:%s@couchdb:5984/" % (app.config['COUCHDB_USER'], app.config['COUCHDB_PASSWORD']))
# Hack to prevent data persisting.
del couchconn['users']

for dbname in ['users', 'posts']:
    if dbname not in couchconn:
        db = couchconn.create(dbname)
    if dbname == 'users':
        # FIXME: Initialize our only user. Use a proper login/password hash method in the future.
        couchconn[dbname].save({"username": "admin", "password": "admin"})


def get_db():
    """Opens a new database connection if there is none yet for the
    current application context.
    """
    couchconn = couchdb.Server(
        "http://%s:%s@couchdb:5984/" % (app.config['COUCHDB_USER'],
                                        app.config['COUCHDB_PASSWORD']))

    return couchconn


@app.route("/")
def index():
    # Fetch posts from database
    couchconn = get_db()
    db = couchconn['posts']
    results = db.view('_all_docs', include_docs=True)

    posts = []
    for row in results.rows:
        posts.append({'id':row.doc.id,'title': row.doc['title'], 'text': row.doc['text'], 'author': row.doc['author'], 'comments': row.doc['comments']})
    app.logger.info(posts)

    return render_template('index.html', posts=posts)


@app.route("/login", methods=['POST', 'GET'])
def login():
    if request.method == "POST":
        app.logger.info(request.form)
        username = request.form.get('username')
        password = request.form.get('password')

        couchconn = get_db()
        db = couchconn['users']

        # Not recommended. use hashed/salted passwords in the future.
        for row in db.find({'selector': {"username": username, "password": password}}):
            # If None not returned, set the user session and
            session['username'] = username
            return redirect(url_for('index'))

        return redirect(url_for('login'))
    else:
        return render_template('login.html')


@app.route("/posts", methods=["GET", "POST"])
def posts():
    if request.method == "GET":
        couchconn = get_db()
        db = couchconn['posts']
        results = db.view('_all_docs', include_docs=True)

        posts = []
        for row in results.rows:
            posts.append(
                {'title': row.doc['title'], 'text': row.doc['text'], 'author': row.doc['author'],
                 'comments': row.doc['comments']})

        return render_template("posts.html", posts=posts)
    if request.method == "POST":
        title = request.form.get('title')
        text = request.form.get('text')
        author = session.get('username')

        couchconn = get_db()
        couchconn['posts'].save({"title": title, "text": text, "author": author, "comments": []})

        return redirect(url_for('index'))

@app.route("/post/<id>", methods=["GET", "POST"])
def get_post(id):
    if request.method == "GET":
        couchconn = get_db()
        postdoc = couchconn['posts'].get(id)
        app.logger.info(postdoc)

        post = {'id':postdoc.id,'title': postdoc['title'], 'text': postdoc['text'],
                'author': postdoc['author'], 'comments': postdoc['comments']}

        return render_template("post.html", post=post)
    if request.method == "POST":
        couchconn = get_db()
        post = couchconn['posts'].get(id)

        author = request.form.get('author')
        text = request.form.get('text')
        post["comments"].append({'author': author, 'text': text})

        couchconn['posts'].save(post)

        return redirect(request.url)

@app.route("/logout")
def logout():
    session.pop('username')
    return redirect(url_for('index'))

if __name__ == "__main__":
    # Only for debugging while developing
    app.run(host="0.0.0.0", debug=True, port=80)
