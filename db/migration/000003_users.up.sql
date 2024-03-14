
CREATE TABLE "users" (
                         "user_id" varchar not null ,
                         "username" varchar PRIMARY KEY  not null ,
                         "password" varchar(255) not null ,
                         "email" varchar(255) not null ,
                          "role_id" serial not null ,
                         "created_at" timestamptz NOT NULL DEFAULT NOW(),
                         "updated_at" timestamptz not null default (now())
);

CREATE TABLE "purchase_requisitions" (
                                         "requisition_id" uuid PRIMARY KEY not null ,
                                         "requester_id" varchar not null ,
                                         "description" varchar not null ,
                                          "title" varchar not null,
                                         "status" varchar(50) not null ,
                                         "date_submitted" timestamptz NOT NULL DEFAULT (now()),
                                         "date_approved" timestamptz not null  default (now()),
                                         "date_rejected" timestamptz not null default (now())
);

CREATE TABLE "quotation_requests" (
                                      "proposal_request_id" bigserial PRIMARY KEY not null ,
                                      "requisition_id" uuid not null ,
                                      "supplier_id" uuid not null ,
                                      "status" varchar(50) not null ,
                                      "date_submitted" timestamptz not null default (now())  ,
                                      "date_approved" timestamptz not null default (now())  ,
                                      "date_rejected" timestamptz not null default (now())
);



CREATE TABLE "suppliers" (
                             "supplier_id" uuid PRIMARY KEY not null ,
                             "supplier_name" varchar(200) not null ,
                             "contact_person" varchar(200) not null ,
                             "contact_mail" varchar(200) not null ,
                             "supplier_type" varchar(200) not null
);

CREATE TYPE "user_type_enum" AS ENUM (
    'procurer',
    'approver',
    'supplier'
    );
--
--
-- CREATE TABLE "user_types" (
--                               "user_type_pk" integer PRIMARY KEY,
--                               "user_type" user_type_enum
-- );

CREATE TABLE "roles"(
    role_id SERIAL PRIMARY KEY,
    role_name user_type_enum NOT NULL UNIQUE
);







ALTER TABLE "purchase_requisitions" ADD FOREIGN KEY ("requester_id") REFERENCES "users" ("username");

ALTER TABLE "quotation_requests" ADD FOREIGN KEY ("requisition_id") REFERENCES "purchase_requisitions" ("requisition_id");

ALTER TABLE "quotation_requests" ADD FOREIGN KEY ("supplier_id") REFERENCES "suppliers" ("supplier_id");

ALTER TABLE "users" ADD FOREIGN KEY ("role_id") REFERENCES "roles" ("role_id");