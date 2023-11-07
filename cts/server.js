const express = require('express');
const bodyParser = require('body-parser');

const app = express();
app.use(bodyParser.json());

// Mock implementation of createOrg function
app.post('/createOrg', (req, res) => {
  // Simulate creating an organization
  const org = {
    id: 'org123',
    name: req.body.name || 'defaultOrgName'
  };
  console.log('Creating organization:', org);
  res.json(org); // Send the mock org object back
});

// Mock sendToAmplitude
app.post('/sendToAmplitude', (req, res) => {
  console.log('Sending data to Amplitude:', req.body);
  res.json({ message: 'Data sent to Amplitude' });
});

// Mock goConvertTestResults
app.post('/goConvertTestResults', (req, res) => {
  console.log('Converting test results:', req.body);
  // Mock conversion here...
  res.json({ message: 'Test results converted' });
});

const port = 3000;
app.listen(port, () => {
  console.log(`Server running on port ${port}`);
});
