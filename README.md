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
4. You can login using credentials username: `admin` and password `admin`.

## Testing
- Manual testing can be done on the running webapp.

## Routes

- `/` fetches the homepage. This returns an index where all posts can be fetched and viewed
- `/post/<id>` fetches a particular post. A post can be viewed here and also commented on.
- `/posts` is for an authenticated user. Here, the authenticated user can create a new blog post.
- `/login` facilitates logging in. For now only the "video game reviewer" can access an account
- `/logout` logs the user out.

## Discussion

## Future Goals
