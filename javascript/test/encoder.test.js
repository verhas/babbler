'use strict';

const test = require('node:test');
const assert = require('node:assert/strict');

const { numberToId } = require('../src/encoder');
const { BLACKLIST, MAX_NUM } = require('../src/constants');

test('numberToId(0) returns a valid 4-syllable id', () => {
  const id = numberToId(0);
  assert.match(id, /^[A-Z][a-z]{3} [A-Z][a-z]{3}$/);
});

test('numberToId(1) differs from numberToId(0)', () => {
  assert.notEqual(numberToId(1), numberToId(0));
});

test('numberToId(10000) matches reference output', () => {
  assert.equal(numberToId(10000), 'Talo Buno');
});

test('numberToId(10001) matches reference output', () => {
  assert.equal(numberToId(10001), 'Patu Luta');
});

test('numberToId(MAX_NUM) succeeds and matches reference output', () => {
  assert.equal(numberToId(MAX_NUM), 'Dobu Zusa');
});

test('numberToId(MAX_NUM + 1) throws', () => {
  assert.throws(() => numberToId(MAX_NUM + 1), RangeError);
});

test('numberToId(-1) throws', () => {
  assert.throws(() => numberToId(-1), RangeError);
});

test('numberToId(1.5) throws', () => {
  assert.throws(() => numberToId(1.5), RangeError);
});

test('capitalization: first letter uppercase only', () => {
  const id = numberToId(42);
  const [w1, w2] = id.split(' ');
  assert.equal(w1, w1[0].toUpperCase() + w1.slice(1).toLowerCase());
  assert.equal(w2, w2[0].toUpperCase() + w2.slice(1).toLowerCase());
});

test('determinism: same input always produces the same output', () => {
  const n = 987654;
  assert.equal(numberToId(n), numberToId(n));
});

test('uniqueness and blacklist-avoidance across the entire valid range', () => {
  const seen = new Set();
  for (let n = 0; n <= MAX_NUM; n += 1) {
    const id = numberToId(n);
    assert.ok(!seen.has(id), `${id} (from ${n}) was already issued for a different number`);
    seen.add(id);

    const [w1, w2] = id.toLowerCase().split(' ');
    assert.ok(!BLACKLIST.has(w1), `${w1} from ${n} is blacklisted`);
    assert.ok(!BLACKLIST.has(w2), `${w2} from ${n} is blacklisted`);
  }
  assert.equal(seen.size, MAX_NUM + 1);
});
