CREATE TABLE IF NOT EXISTS "authors" (
    "id" VARCHAR(36) PRIMARY KEY,
    "name" VARCHAR(255) NOT NULL,
    "bio" VARCHAR(255) NOT NULL,
    "created_at" TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    "updated_at" TIMESTAMP WITH TIME ZONE NULL
);

CREATE TABLE IF NOT EXISTS "subjects" (
    "id" VARCHAR(36) PRIMARY KEY,
    "name" VARCHAR(255) NOT NULL,
    "description" VARCHAR(255) NOT NULL,
    "created_at" TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    "updated_at" TIMESTAMP WITH TIME ZONE NULL
);

CREATE TABLE IF NOT EXISTS "books" (
    "id" VARCHAR(36) PRIMARY KEY,
    "title" VARCHAR(255) NOT NULL,
    "isbn" VARCHAR(13) NOT NULL,
    "year" INT NOT NULL,
    "pages" INT NOT NULL,
    "created_at" TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    "updated_at" TIMESTAMP WITH TIME ZONE NULL
);

CREATE TABLE IF NOT EXISTS "books_authors" (
    "book_id" VARCHAR(36) NOT NULL,
    "author_id" VARCHAR(36) NOT NULL,
    PRIMARY KEY ("book_id", "author_id"),
    FOREIGN KEY ("book_id") REFERENCES "books" ("id"),
    FOREIGN KEY ("author_id") REFERENCES "authors" ("id")
);

CREATE TABLE IF NOT EXISTS "books_subjects" (
    "book_id" VARCHAR(36) NOT NULL,
    "subject_id" VARCHAR(36) NOT NULL,
    PRIMARY KEY ("book_id", "subject_id"),
    FOREIGN KEY ("book_id") REFERENCES "books" ("id"),
    FOREIGN KEY ("subject_id") REFERENCES "subjects" ("id")
);

CREATE TABLE IF NOT EXISTS "quotes" (
    "id" VARCHAR(36) PRIMARY KEY,
    "content" VARCHAR(255) NOT NULL,
    "page" INT NOT NULL,
    "book_id" VARCHAR(36) NOT NULL,
    "created_at" TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    "updated_at" TIMESTAMP WITH TIME ZONE NULL,
    FOREIGN KEY ("book_id") REFERENCES "books" ("id")
);

