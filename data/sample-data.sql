--
-- Sample data for groups
--

INSERT INTO groups(name) VALUES
("Trial"),
("Basic"),
("Premium");

--
-- Sample data for images
--

INSERT INTO images(width, height, url, title, artist, gallery, organization, media) VALUES
(512, 362, "https://lh3.ggpht.com/iqWhX4gA_fDh_XUFb6bjoCVusT6lIAW4r1KCbBFLDyeqSczEwdzzwEA4SUn5Z_qDlef4", "Bodegón", "Rubén Darío Velázquez", NULL, "Google Open Gallery", "Art"),
(379, 512, "https://lh3.ggpht.com/RbL9u5vdZMxltINmK5w48iS3_fMhfjVR7SxD-MMB1OZShj9OnynDhdAcvk7N3JDzvdbR-tk", "Portrait of Gabriel Bernard de Rieux", "Maurice-Quentin de La Tour", "The J. Paul Getty Museum", "Google Cultural Institute", "Art"),
(455, 512, "https://lh3.ggpht.com/tF611zF3LcdZLnJqCPmnNwvsuN4rAjwPNzrO5Uu4G8HT1iPUzmBI7w", "The wattles", "Elioth Gruner", "Art Gallery of New South Wales", "Google Cultural Institute", "Art"),
(512, 384, "https://lh4.ggpht.com/CpBfF0vqJ2eqiL5Af3J5VtbqmIZQ9IpM2NJFdmE_49gGV2YIm0muJnMibzjpm-9cqxfOgg", "Envol pastel", "Octave Pixel", "galerie TACT: Fresh artistes", "Google Open Gallery", "Art"),
(2048, 1365, "https://lh3.ggpht.com/XprZKBFc-ErF-K_KKJEcWqiwwrsbn3BCg4yBlManYe0x2c-E6K4m2EFT1vYjrrae_VUMslw", "Manhattan", "Anatoliy Urbanskiy", "Featured photos", "Google+", "Photo"),
(512, 222, "https://lh3.ggpht.com/6swp7Zh21xlvw8NOEXYrMrUrcMSzzNvPd-3nZsR5RrrTQRTaMrKsdhIx4nHbVYyNmib2Jo_T9rR-gXQ", NULL, "Grzegorz Głowaty", NULL, NULL, "Photo"),
(512, 342, "https://lh3.ggpht.com/Vor04AFDXMTOb9tE0D0eK5zkuyKwHjvfjegcFd1sJxlfZWMeTFbjuMTo0PPQ6pHHQx5--Cc4eT48d9g", NULL, "Raja Ramakrishnan", NULL, NULL, "Photo"),
(512, 311, "https://lh3.ggpht.com/sDM3YrHDSPg4o75K8cna0M2N0yMSkXEC80Cc8AQL0YiNntHbmo9YPkRJ936wmsDubXGeBCxneBeTFpU", NULL, "Paul Moody", NULL, NULL, "Photo"),
(2048, 1365, "https://lh3.ggpht.com/pwOiNl9yVpvleSguShs3MtXwUOPfMLtrQ3B4WUEoj0UR2bXgUEG1", "Early morning Bronte Beach", "Gil Baillache", NULL, NULL, "Photo");

--
-- Sample data for images-groups
--

INSERT INTO images_groups(group_id, image_id) VALUES
(1, 5),
(2, 1),
(2, 4),
(2, 5),
(3, 1),
(3, 2),
(3, 3),
(3, 4),
(3, 5),
(3, 6),
(3, 7),
(3, 8),
(3, 9);

--
-- Sample data for users
--

INSERT INTO users (group_id, guid, first_name, middle_init, last_name, email, addr, city, state, zip, gender, status) VALUES
(2, "3efd7840-14f8-4c34-83fe-8b036661e607", "Charles", "A", "Bohner", "CharlesABohner@teleworm.us", "4347 Walnut Avenue", "Boise", "ID", "83702", "M", TRUE),
(2, "6be47a14-1e5c-422a-8b41-8d95e469002f", "Julie", "C", "Roesch", "JulieCRoesch@rhyta.com", "1971 Tyler Avenue", "Ft Lauderdale", "FL", "33311", "F", TRUE),
(2, "568e1607-752d-4853-ae21-a1b29d3359f6", "Bob", "C", "Smith", "BobCSmith@rhyta.com", "320 Barfield Lane", "Indianapolis", "IN", "46278", "M", TRUE),
(2, "6c672ab8-c3d4-4f7b-81d1-b0a11bd7b984", "Dorothy", "J", "Fernandez", "DorothyJFernandez@armyspy.com", "2512 Douglas Dairy Road", "Gate City", "VA", "24251", "F", TRUE),
(3, "7db0f8e6-4d65-4df2-b3ed-de6b8e02dac4", "William", "G", "Lukes", "WilliamGLukes@dayrep.com", "1763 Brookside Drive", "Birmingham", "AL", "35222", "M", TRUE),
(2, "dcadc958-c1c9-44e7-ba79-91ba10b2d05c", "Anna", "H", "Khan", "AnnaHKhan@rhyta.com", "2580 Williams Mine Road", "Somerville", "NJ", "08876", "F", TRUE),
(2, "d80fbec2-0ce3-40cb-9f07-70e236372d6c", "Randall", "K", "Fellows", "RandallKFellows@jourrapide.com", "2447 Emma Street", "Lubbock", "TX", "79410", "M", TRUE),
(2, "7017b2e1-8d87-4dfd-87d6-5579b82741b4", "Kimberly", "M", "Miller", "KimberlyMMiller@armyspy.com", "2329 Chardonnay Drive", "La Center", "WA", "98629", "F", TRUE),
(2, "26d52f8c-49b7-426e-931d-8ad5f351e4fa", "Bill", "T", "Rios", "BillTRios@jourrapide.com", "3823 Eagle Street", "Fairview Heights", "IL", "62208", "M", TRUE),
(3, "3f88265e-b2b2-450c-878d-59f14ee76bfd", "Jennifer", "C", "Taylor", "JenniferCTaylor@armyspy.com", "4771 Simons Hollow Road", "Middletown", "PA", "10940", "F", TRUE),
(3, "3695ea56-b8c4-47f0-8e70-8825e3cf81d9", "John", "R", "Breen", "JohnRBreen@dayrep.com", "4530 Collins Street", "Tampa", "FL", "33624", "M", TRUE),
(3, "6c5ef621-4d48-4b65-9199-3be4e8431039", "Zada", "R", "Coppa", "ZadaRCoppa@armyspy.com", "2671 Tree Frog Lane", "Kansas City", "MO", "64106", "F", TRUE),
(2, "a6e8d104-7f47-417c-af0b-e4ebd96b15ae", "Martin", "A", "Walker", "MartinAWalker@armyspy.com", "2938 Kimberly Way", "Boston", "MA", "02109", "M", TRUE),
(2, "95193d50-2652-407d-86f1-2121709f5a32", "Margaret", "D", "Lovell", "MargaretDLovell@jourrapide.com", "3469 Wildrose Lane", "Detroit", "MI", "48226", "F", TRUE),
(2, "37e02d57-1594-4815-a25e-ed6ca840adde", "Leroy", "P", "Keith", "LeroyPKeith@teleworm.us", "3027 Hickory Heights Drive", "Linthicum Heights", "MD", "21090", "M", FALSE),
(2, "5f2156e6-d041-4112-9acd-5f5e2f6517b9", "Dawna", "M", "Duvall", "DawnaMDuvall@rhyta.com", "4122 Candlelight Drive", "Baytown", "TX", "77520", "F", TRUE),
(2, "736c9542-6858-4cab-9216-bbbdc7908def", "Eric", "W", "Johnson", "EricWJohnson@teleworm.us", "3665 North Street", "Charlottesville", "VA", "22903", "M", TRUE),
(1, "35cdb003-825b-4e76-a660-2e1102ecfa9a", "Helen", "P", "Bonds", "HelenPBonds@rhyta.com", "3428 Heavner Avenue", "Carrollton", "GA", "30117", "F", TRUE),
(2, "58b3af8f-b657-4469-90f7-5fd17ad3c4fa", "Glen", "L", "Harper", "GlenLHarper@dayrep.com", "131 Adamsville Road", "Laredo", "TX", "78040", "M", TRUE),
(2, "5e8b18b3-e2ad-4108-b67e-92a4f1533acb", "Ruby", "M", "Mendez", "RubyMMendez@armyspy.com", "1853 Davis Court", "Stlouis", "IL", "63101", "F", TRUE),
(2, "9c232198-4ae3-43b9-80f9-112b0e3046f5", "Edward", "M", "Orr", "EdwardMOrr@jourrapide.com", "2045 Havanna Street", "Greensboro", "NC", "27407", "M", TRUE),
(2, "2772827a-cf9c-4774-a571-4e7b94c966e5", "Judith", "R", "Moser", "JudithRMoser@armyspy.com", "3007 Brown Avenue", "Knoxville", "TN", "37902", "F", TRUE),
(3, "a33347ad-e8b7-4456-832d-0f9cbed9be51", "Philip", "M", "Heckman", "PhilipMHeckman@armyspy.com", "312 Jennifer Lane", "Creedmoor", "NC", "27522", "M", TRUE),
(2, "d9b8bdc0-d580-4988-8725-37d823bd016e", "Charlotte", "J", "Salgado", "CharlotteJSalgado@dayrep.com", "1631 Armbrester Drive", "West Los Angeles", "CA", "90025", "F", TRUE),
(2, "e2da0a04-327d-4ee1-bc36-fddf4c489207", "Noel", "B", "Roybal", "NoelBRoybal@rhyta.com", "4959 Bridge Avenue", "Hayes", "LA", "70646", "M", FALSE),
(2, "e8ae1ec2-dbcb-424c-9072-da0eaa3d5c30", "Pansy", "B", "Robbins", "PansyBRobbins@teleworm.us", "4033 Elm Drive", "New York", "NY", "10011", "F", TRUE),
(2, "b15ba298-6493-4363-979d-600d2672aeb6", "Michael", "A", "Knapp", "MichaelAKnapp@jourrapide.com", "3740 Waldeck Street", "Grapevine", "TX", "76051", "M", TRUE),
(2, "895c7e8b-011d-4399-a3f2-85848b584b0a", "Lois", "N", "Smith", "LoisNSmith@teleworm.us", "770 Geneva Street", "Bellerose", "NY", "11426", "F", TRUE),
(2, "f7f3b7f5-ee68-401e-9e3f-64541297704c", "Richard", "J", "Sawyer", "RichardJSawyer@rhyta.com", "3176 Dancing Dove Lane", "New York", "NY", "10013", "M", TRUE),
(3, "ce8372e1-3d70-4851-9a45-1debdb619043", "Maude", "P", "Hansen", "MaudePHansen@armyspy.com", "2691 Beechwood Drive", "Pittsburgh", "PA", "15219", "F", TRUE),
(3, "f2c1957e-8702-49b5-b7e1-09cff665950f", "Lawrence", "M", "Staudt", "LawrenceMStaudt@dayrep.com", "4480 Jones Street", "Grapevine", "TX", "76051", "M", TRUE),
(2, "cb85afb0-af18-49fb-93ce-d17cb7b7c4f0", "Jennifer", "R", "Stclair", "JenniferRStclair@jourrapide.com", "3654 Arbutus Drive", "Miami", "FL", "33179", "F", TRUE),
(1, "250fe497-4260-4331-87c4-a582fcdd88bf", "Kyle", "M", "Allen", "KyleMAllen@dayrep.com", "4084 Whaley Lane", "Milwaukee", "WI", "53212", "M", TRUE),
(2, "948b584a-1d74-4e3b-9d57-81e13c0a6149", "Amy", "R", "Stump", "AmyRStump@rhyta.com", "4239 Rhode Island Avenue", "Washington", "DC", "20008", "F", TRUE),
(2, "fb5a8306-6821-497c-bf32-86b2deb92264", "Richard", "E", "Dean", "RichardEDean@dayrep.com", "1680 Bagwell Avenue", "Eustis", "FL", "32726", "M", TRUE),
(2, "880bb6b6-a2a7-4953-8c15-653d39082a60", "Pamela", "J", "Williams", "PamelaJWilliams@teleworm.us", "2093 Browning Lane", "Oxford", "NY", "13830", "F", TRUE),
(3, "3e7aee4c-1736-43f7-9954-598165340b35", "Sean", "J", "Stage", "SeanJStage@armyspy.com", "3879 Giraffe Hill Drive", "Dallas", "TX", "75204", "M", FALSE),
(2, "d8828793-6406-4374-aa72-16d205b9332d", "Marilyn", "V", "Coleman", "MarilynVColeman@teleworm.us", "3076 Fancher Drive", "Dallas", "TX", "75225", "F", TRUE),
(3, "3558782a-75f4-4191-b941-889e113f8607", "Eric", "M", "Ramey", "EricMRamey@dayrep.com", "2947 Roguski Road", "Natchitoches", "LA", "71457", "M", TRUE),
(1, "59fc87e1-08d3-4697-8c5d-7b2f8091382f", "Dianne", "J", "Williams", "DianneJWilliams@dayrep.com", "4197 Kincheloe Road", "Gresham", "OR", "97030", "F", TRUE);
