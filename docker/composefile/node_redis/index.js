    const express = require('express');
    const = require('redis');
    const app = express();
    
    const client = redis.createClient({
        host: ''
    });
    client.set('visits', 0);
    
    app.get('/', (req, res) => {
      res.send('How are you doing');
    });
     
    app.listen(8080, () => {
      console.log('Listening on port 8080');
    });
