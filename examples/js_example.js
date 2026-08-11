// Run from the repo root: node examples/js_example.js
'use strict';

const { numberToId } = require('../javascript/src');

// Typical usage: give a friendly display name to each row in an
// auto-increment sequence (e.g. a database primary key).
for (let userId = 0; userId < 5; userId += 1) {
  console.log(`user #${userId} -> ${numberToId(userId)}`);
}
