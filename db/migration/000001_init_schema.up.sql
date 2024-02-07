
CREATE TYPE UserTypeENUM AS ENUM (
    'procurer',
    'approver',
    'supplier'
    );



CREATE TABLE "UserTypes" (
                             "UserTypePK" integer,
                             "UserType" UserTypeENUM
);




CREATE TABLE "users" (
                         "userid"  integer not null ,
                         "username" varchar not null ,
                         "password" varchar not null ,
                         "email" varchar not null ,
                         "usertypefk" integer not null,
                        "created_at" timestamptz not null ,
                        "updated_at" timestamptz not null


);


ALTER TABLE "UserTypes"   ADD CONSTRAINT user_type_pk_unique UNIQUE("UserTypePK");



CREATE TABLE "purchaserequisitions" (
                                        "requisitionid" varchar PRIMARY KEY ,
                                        "requesterid" varchar not null ,
                                         "title" varchar not null ,
                                        "description" varchar not null ,
                                        "status" varchar not null ,
                                        "datesubmitted" timestamptz not null default (now()),
                                        "dateapproved" timestamptz not null  default (now()),
                                        "daterejected" timestamptz not null  default (now())
);

CREATE TABLE "QuotationRequests" (
                                     "ProposalRequestID" integer,
                                     "RequisitionID" varchar,
                                     "VendorID" integer,
                                     "Status" varchar,
                                     "DateSubmitted" timestamptz not null default (now()),
                                     "DateApproved" timestamptz default (now()),
                                     "DateRejected" timestamptz default (now())
);

CREATE TABLE "Vendors" (
                           "VendorID" integer,
                           "VendorName" varchar,
                           "ContactPerson" varchar,
                           "ContactEmail" varchar,
                           "VendorType" varchar
);

ALTER TABLE "users" ADD FOREIGN KEY ("usertypefk") REFERENCES "UserTypes" ("UserTypePK");

ALTER TABLE "purchaserequisitions" ADD FOREIGN KEY ("requesterid") REFERENCES "users" ("userid");

ALTER TABLE "QuotationRequests" ADD FOREIGN KEY ("RequisitionID") REFERENCES "purchaserequisitions" ("requisitionid");

ALTER TABLE "QuotationRequests" ADD FOREIGN KEY ("VendorID") REFERENCES "Vendors" ("VendorID");
