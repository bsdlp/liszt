/**
 * @flow
 * @relayHash f2872e628eab9bcf0ae14d17dbe38762
 */

/* eslint-disable */

'use strict';

/*::
import type {ConcreteBatch} from 'relay-runtime';
export type AppAllBuildingQueryResponse = {|
  +Buildings: $ReadOnlyArray<?{|
    +id: string;
    +name: string;
  |}>;
|};
*/


/*
query AppAllBuildingQuery {
  Buildings {
    id
    name
  }
}
*/

const batch /*: ConcreteBatch*/ = {
  "fragment": {
    "argumentDefinitions": [],
    "kind": "Fragment",
    "metadata": null,
    "name": "AppAllBuildingQuery",
    "selections": [
      {
        "kind": "LinkedField",
        "alias": null,
        "args": null,
        "concreteType": "Building",
        "name": "Buildings",
        "plural": true,
        "selections": [
          {
            "kind": "ScalarField",
            "alias": null,
            "args": null,
            "name": "id",
            "storageKey": null
          },
          {
            "kind": "ScalarField",
            "alias": null,
            "args": null,
            "name": "name",
            "storageKey": null
          }
        ],
        "storageKey": null
      }
    ],
    "type": "Query"
  },
  "id": null,
  "kind": "Batch",
  "metadata": {},
  "name": "AppAllBuildingQuery",
  "query": {
    "argumentDefinitions": [],
    "kind": "Root",
    "name": "AppAllBuildingQuery",
    "operation": "query",
    "selections": [
      {
        "kind": "LinkedField",
        "alias": null,
        "args": null,
        "concreteType": "Building",
        "name": "Buildings",
        "plural": true,
        "selections": [
          {
            "kind": "ScalarField",
            "alias": null,
            "args": null,
            "name": "id",
            "storageKey": null
          },
          {
            "kind": "ScalarField",
            "alias": null,
            "args": null,
            "name": "name",
            "storageKey": null
          }
        ],
        "storageKey": null
      }
    ]
  },
  "text": "query AppAllBuildingQuery {\n  Buildings {\n    id\n    name\n  }\n}\n"
};

module.exports = batch;
