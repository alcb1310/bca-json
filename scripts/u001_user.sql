INSERT INTO company (id, ruc, name, employees, is_active)
VALUES ('4b114c26-b038-4cfa-ae6e-ad46c73ef59d', '1704749652001', 'Test Company', 3, true);

INSERT INTO "user" (id, email, name, password, company_id, role_id)
VALUES ('0cd002ff-2c33-460b-8876-73e51dfb053e', 'test@test.com', 'Test User', '$2a$14$JVprUuqit3yVubTrsm92lemqP7sbeKRP8wwDnsyDRWQceEcPJMw2e', '4b114c26-b038-4cfa-ae6e-ad46c73ef59d', 'a');

INSERT INTO "user" (id, email, name, password, company_id, role_id)
VALUES ('0cd001ff-2c33-460b-8876-73e51dfb053e', 'a@b.c', 'Another User', '$2a$14$JVprUuqit3yVubTrsm92lemqP7sbeKRP8wwDnsyDRWQceEcPJMw2e', '4b114c26-b038-4cfa-ae6e-ad46c73ef59d', 'a');

INSERT INTO company (id, ruc, name, employees, is_active)
VALUES ('4b214c26-b038-4cfa-ae6e-ad46c73ef59d', '1791838300001', 'Another Company', 3, true);

INSERT INTO "user" (id, email, name, password, company_id, role_id)
VALUES ('0cd002ef-2c33-460b-8876-73e51dfb053e', 'unknown@unknown.com', 'Unknown user', 't', '4b214c26-b038-4cfa-ae6e-ad46c73ef59d', 'a');
