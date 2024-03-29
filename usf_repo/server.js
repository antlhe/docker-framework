const express = require('express');
const bodyParser = require('body-parser');
const xmlparser = require('express-xml-bodyparser');

const app = express();
app.use(bodyParser.json());
app.use(bodyParser.text({ type: 'text/xml' }));
app.use(xmlparser());

app.post('/createOrg', (req, res) => {
  const org = {
    id: 'org123',
    name: req.body.name || 'defaultOrgName'
  };
  console.log('Creating organization:', org);
  res.json(org);
});

app.post('/sendToAmplitude', (req, res) => {
  if (req.headers['content-type'] === 'text/xml' || req.headers['content-type'] === 'application/xml') {
    console.log('Receiving XML data...');
    console.log('Sending data to Amplitude:', req.body);

    res.status(200).send({ message: 'XML data sent to Amplitude' });
  } else {
    res.status(400).send({ error: 'Invalid content type. This endpoint only accepts XML data.' });
  }
});


const port = 3000;
app.listen(port, () => {
  console.log(`Server running on port ${port}`);
});
