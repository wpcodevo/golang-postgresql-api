CREATE TABLE
    "posts" (
        "id" UUID NOT NULL DEFAULT (uuid_generate_v4()),
        "title" VARCHAR NOT NULL,
        "category" VARCHAR NOT NULL,
        "content" VARCHAR NOT NULL,
        "image" VARCHAR NOT NULL,
        "created_at" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
        "updated_at" TIMESTAMP(3) NOT NULL,
        CONSTRAINT "posts_pkey" PRIMARY KEY ("id")
    );

CREATE UNIQUE INDEX "posts_title_key" ON "posts"("title");