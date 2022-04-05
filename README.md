# Intro

Server side of another pet store project. This project intention is to learn go language and get more familiar with postgres db.

# Build and run

To run this project you need to have installed go dependencies and docker. The current working branch is **develop**.

1. Checkout this project to your working directory
2. Create a **.env** file in root of the project
3. Fill required parameters like db_user, db_pass and scheme, **note** that default db user is **supercat**, it can be changed in **docker-compose.yml**:

+ DB_URL="postgres://db_user_name:_db_user_pass_@localhost:5432/_db_scheme_?sslmode=disable"

+ DB_PASS=db_pass

5. Export variables to your environment: **export $(cat .env | xargs)**
6. Run db migrations like: **sh run_migrations.sh**
7. Optional: the random data may be generated using psql command line inside db container, container name is displayed via **docker ps** command: **docker exec -it {container_name} psql**
8. All things are set now. The project can be run: **go run .**

If there is problem with dependencies **go mod tidy** can help


# Features and improvements

+ Proper error handling, without stop the entire application
+ Add DI framework
+ Make pet store an actual pet store, not just cat-store - e.g. add other animals
+ Recheck JWT security
+ Do not import user service in controller, it must be under auth service or vice versa
+ Tests...
+ More things to come