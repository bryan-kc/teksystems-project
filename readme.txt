## Context ##

This is a take home coding test to demonstrate capabilities in the development and operations space.
The task is described in the user story above and acceptance criteria below. Implementation details such as language, framework and API contract are up to the engineer.
The code should be shared online in a way which allows access to a validating employee (ideally in a git repository).
Instructions should be provided on how to build and run the application.
    It will be built, run and validated on the workstation of the test validator (so no shared deployment or code is necessary).
The test validator will be another software engineer and will look into the functional code, the build system and everything else.
For the purposes of this test the implementation should be scoped to an example back end API that would enable the desired story. Commit and share whatever you have at the end of the allotted timeframe, don't worry if you didn't fully complete the task.

## Acceptance Criteria ##

A blog post will show a title, article text (plain text) and an author name
Comments are made on blog posts and show comment text (plain text) and an author name

## Non-functional Requirements ##

The code should be the type of code you would consider production ready and could be supported by a team. Write the sort of code you would want to receive from another engineer.
The application must have a build system
The application build should be built or compiled in a docker container, so the build is portable
The application build should produce a docker container as an artifact, so the deployment is portable
The application should not have any runtime dependencies (so the datastore needs to be in memory or similar

