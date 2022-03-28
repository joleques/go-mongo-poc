db = db.getSiblingDB('golang-test');

db.createUser({user:"golang-test",pwd:"test123","roles":[{"role":"readWrite","db":"golang-test"},{"role":"dbAdmin","db":"golang-test"}]});
