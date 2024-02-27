CREATE TABLE "business_particulars" (
                                        "reg_business_name" varchar not null ,
                                        "brela_reg_number" varchar PRIMARY KEY not null ,
                                        "user_id" serial not null ,
                                        "po_box" varchar not null ,
                                        "occupation_location" varchar not null ,
                                        "country" varchar(255) not null ,
                                        "region" varchar(255) not null ,
                                        "tin" varchar(255) not null ,
                                        "first_name" varchar(255) not null ,
                                        "last_name" varchar(255) not null ,
                                        "contact" varchar not null ,
                                        "website" varchar not null

);

