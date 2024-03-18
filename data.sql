INSERT INTO actor (id, name, gender, date_birthday)
VALUES (
        1,
        'alexey',
        'male',
        CURRENT_TIMESTAMP
    );
INSERT INTO actor (id, name, gender, date_birthday)
VALUES (
        2,
        'jenya',
        'female',
        CURRENT_TIMESTAMP
    );
INSERT INTO films (id, title, description, date_premiere, rating)
VALUES (
        1,
        'film1',
        'good film',
        2000,
        7
    );
INSERT INTO films (id, title, description, date_premiere, rating)
VALUES (
        2,
        'film2',
        'so so film',
        2002,
        4
    );
INSERT INTO actors_films (film_id, actor_id)
VALUES (1, 1);
INSERT INTO actors_films (film_id, actor_id)
VALUES (1, 2);
INSERT INTO actors_films (film_id, actor_id)
VALUES (2, 1);
INSERT INTO actors_films (film_id, actor_id)
VALUES (2, 2);
insert into users
values (1, 'useruser', 'password1', 'isUser'),
    (2, 'useradmin', 'password2', 'isAdmin');