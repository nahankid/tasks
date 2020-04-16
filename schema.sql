BEGIN TRANSACTION;
    CREATE TABLE users (
        id SERIAL primary key,
        username varchar(100),
        password varchar(1000),
        email varchar(100)
    );
    INSERT INTO users VALUES(1,'suraj','suraj','sapatil@live.com');


    --category
    CREATE TABLE category( 
        id SERIAL primary key,
        name varchar(1000) not null, 
        user_id INTEGER references users(id)
    );

    INSERT INTO "category" VALUES(1,'TaskApp',1);

    --status
    BEGIN TRANSACTION;
    CREATE TABLE status (
        id SERIAL primary key,
        status varchar(50) not null
    );
    INSERT INTO "status" VALUES(1,'COMPLETE');
    INSERT INTO "status" VALUES(2,'PENDING');
    INSERT INTO "status" VALUES(3,'DELETED');
    COMMIT;

    --task

    BEGIN TRANSACTION;
    CREATE TABLE task (
        id SERIAL primary key,
        title varchar(100),
        content text,
        created_date timestamp,
        last_modified_at timestamp,
        finish_date timestamp,
        priority integer, 
        cat_id INTEGER references category(id), 
        task_status_id INTEGER references status(id), 
        due_date timestamp, 
        user_id INTEGER references users(id), 
        hide int
    );

    INSERT INTO "task" VALUES(1,'Publish on github','Publish the source of tasks and picsort on github',NOW(),NOW(),NULL,3,1,1,NULL,1,0);
    INSERT INTO "task" VALUES(4,'gofmtall','The idea is to run gofmt -w file.go on every go file in the listing, *Edit turns out this is is difficult to do in golang **Edit barely 3 line bash script. ',NOW(),NOW(),NULL,3,1,1,NULL,1,0);

    CREATE TABLE comments(
        id SERIAL primary key,
        content text,
        taskID INTEGER references task(id),
        created timestamp,
        user_id INTEGER references users(id)
    );

    CREATE TABLE files(name varchar(1000) not null, autoName varchar(255) not null, user_id INTEGER references users(id), created_date timestamp);
    COMMIT;