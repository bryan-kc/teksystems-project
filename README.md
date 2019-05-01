# TEKsystems DevOps Project

## Technologies used:

- Based on https://github.com/tiangolo/uwsgi-nginx-flask-docker/blob/master/python3.7/Dockerfile
- CouchDB for storage of posts and users.

## User Story
As an avid video game reviewer
I want a way to create blog posts for my video game reviews
So that I can share my reviews in a way that my readers can respond to

### Breakdown
- Single user (video game reviewer)
- Way to create blog posts for video game reviews
    - What is a blog post?
    - What makes up a 'video game review?'
        - Title, text discussion and author
- So that I can share my reviews in a way that my readers can respond to
    - _Share my reviews_
        - Sharing is inherent with post creation
    - "In a way my readers can respond to"
        - Implies a comment section

### Acceptance Criteria
- A blog post will show a title, article text (plain text) and an author name
- Comments are made on blog posts and show comment text (plain text) and an author name

## Usage
- In the root of the project run:
1. `docker-compose build` to build the webapp image
2. `docker-compose up` to run the stack.
3. In your browser access `0.0.0.0:80` or `localhost:80`
    1. You may have to refresh the page once if the app crashes as the webapp isn't ready on first load.
4. You can login using credentials username: `admin` and password `admin`.
    1. As a logged in user, you can make posts under the `posts` tab in the header

5. Either as a logged in or logged out user, you can click created blog posts to be 
taken to the post itself, where you can comment with a name and a comment.
6. Values are stored under the CouchDB database in the `users` and `posts` databases specifically. 
    1. `users` have a username and a password
    2. `posts` have an author, text and comments section which is a list of `author` and `text` 
    storing the name and the text content of the comment.
7. You can access couchDB to see the status of the documents themselves at `0.0.0.0:5984/_utils` 
with user `admin` password: `password`


## Testing
- Manual testing can be done on the running webapp.

## Routes

- `/` fetches the homepage. This returns an index where all posts can be fetched and viewed
- `/post/<id>` fetches a particular post. A post can be viewed here and also commented on.
- `/posts` is for an authenticated user. Here, the authenticated user can create a new blog post.
- `/login` facilitates logging in. For now only the "video game reviewer" can access an account
- `/logout` logs the user out.

## Discussion
- The application 

## Future Goals
- Add tests
- Utilize a more rigorous object relationship for the contents being stored
- Decouple comments from the posts themselves, and have each comment link to a postID
    - Comments can also store timestamps.
    - Comments could also have replies.
- Update UI
- Fix first start crash bug
- Add proper authentication. 
    - Add user roles such as admins/editors and commenters who will have accounts but not be able to create reviews
- Posts could have a score field. Either a rating system (stars) or a score out of 100.
- Users could have profile pictures
