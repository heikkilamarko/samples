CREATE TABLE DEMO.DEMO_EVENT (
    EVENT_ID   NUMBER GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
    EVENT_TYPE VARCHAR2(255) NOT NULL,
    EVENT_DATA CLOB,
    CREATED_AT TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,
    CREATED_BY VARCHAR2(255) NOT NULL
);
