package server

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Database struct {
	DB *sqlx.DB
}

func SetupDatabase() Database {
	// Connect to the database
	db, err := sqlx.Connect("postgres", "user=postgres password=root dbname=doalivros sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}

	// Run the setup SQL commands
	setupSQL := `
		CREATE TABLE IF NOT EXISTS public.users (
		    id serial4 NOT NULL,
		    first_name varchar(255) NOT NULL,
		    last_name varchar(255) NOT NULL,
		    email varchar(255) NOT NULL,
		    "password" varchar(255) NOT NULL,
		    CONSTRAINT users_email_key UNIQUE (email),
		    CONSTRAINT users_pkey PRIMARY KEY (id)
		);

		CREATE TABLE IF NOT EXISTS public.books (
		    id serial4 NOT NULL,
		    title varchar(255) NOT NULL,
		    author varchar(255) NOT NULL,
		    user_id int8 NULL,
		    donating bool DEFAULT false NULL,
		    CONSTRAINT books_pkey PRIMARY KEY (id)
		);

		CREATE TABLE IF NOT EXISTS public.donated_books (
		    id serial4 NOT NULL,
		    from_user_id int8 NOT NULL,
		    to_user_id int8 NOT NULL,
		    book_id int8 NOT NULL,
		    to_user_name varchar(255) NULL,
		    CONSTRAINT donated_books_pkey PRIMARY KEY (id)
		);

		DO $$ BEGIN
		    IF NOT EXISTS (SELECT 1 FROM information_schema.table_constraints WHERE constraint_name = 'donated_books_book_id_fkey') THEN
		        ALTER TABLE public.donated_books 
		        ADD CONSTRAINT donated_books_book_id_fkey FOREIGN KEY (book_id) 
		        REFERENCES public.books(id) ON DELETE CASCADE;
		    END IF;
		END $$;

		DO $$ BEGIN
		    IF NOT EXISTS (SELECT 1 FROM information_schema.table_constraints WHERE constraint_name = 'donated_books_from_user_id_fkey') THEN
		        ALTER TABLE public.donated_books 
		        ADD CONSTRAINT donated_books_from_user_id_fkey FOREIGN KEY (from_user_id) 
		        REFERENCES public.users(id) ON DELETE CASCADE;
		    END IF;
		END $$;

		DO $$ BEGIN
		    IF NOT EXISTS (SELECT 1 FROM information_schema.table_constraints WHERE constraint_name = 'donated_books_to_user_id_fkey') THEN
		        ALTER TABLE public.donated_books 
		        ADD CONSTRAINT donated_books_to_user_id_fkey FOREIGN KEY (to_user_id) 
		        REFERENCES public.users(id) ON DELETE CASCADE;
		    END IF;
		END $$;
	`

	// Execute the SQL statements
	_, err = db.Exec(setupSQL)
	if err != nil {
		log.Fatalln(err)
	}

	// Return the database connection
	return Database{
		DB: db,
	}
}
