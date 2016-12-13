var mysql = require('mysql');

{
  host     : 'localhost',
  user     : 'root',
  password : 'exoticpi',
  database : 'exoticars',
}
);

connection.connect();

// Take new quantity and multiply it with price. Save it as new TotalPrice
 
var queryString = 'INSERT * TO Cart';

connection.query(queryString, function(err, rows, fields) {
    if (err) throw err;

    for (var i in rows) {
        console.log('Post Titles: ', rows[i].idProducts);
    }
});

connection.end();
