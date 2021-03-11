/* tslint:disable */
/* eslint-disable */
// @ts-nocheck

import { ConcreteRequest } from "relay-runtime";
export type UserListAllUsersQueryVariables = {};
export type UserListAllUsersQueryResponse = {
    readonly users: {
        readonly edges: ReadonlyArray<{
            readonly node: {
                readonly id: string;
                readonly username: string;
            };
        } | null> | null;
    } | null;
};
export type UserListAllUsersQuery = {
    readonly response: UserListAllUsersQueryResponse;
    readonly variables: UserListAllUsersQueryVariables;
};



/*
query UserListAllUsersQuery {
  users {
    edges {
      node {
        id
        username
      }
    }
  }
}
*/

const node: ConcreteRequest = (function(){
var v0 = [
  {
    "alias": null,
    "args": null,
    "concreteType": "UserConnection",
    "kind": "LinkedField",
    "name": "users",
    "plural": false,
    "selections": [
      {
        "alias": null,
        "args": null,
        "concreteType": "UserEdge",
        "kind": "LinkedField",
        "name": "edges",
        "plural": true,
        "selections": [
          {
            "alias": null,
            "args": null,
            "concreteType": "User",
            "kind": "LinkedField",
            "name": "node",
            "plural": false,
            "selections": [
              {
                "alias": null,
                "args": null,
                "kind": "ScalarField",
                "name": "id",
                "storageKey": null
              },
              {
                "alias": null,
                "args": null,
                "kind": "ScalarField",
                "name": "username",
                "storageKey": null
              }
            ],
            "storageKey": null
          }
        ],
        "storageKey": null
      }
    ],
    "storageKey": null
  }
];
return {
  "fragment": {
    "argumentDefinitions": [],
    "kind": "Fragment",
    "metadata": null,
    "name": "UserListAllUsersQuery",
    "selections": (v0/*: any*/),
    "type": "Query",
    "abstractKey": null
  },
  "kind": "Request",
  "operation": {
    "argumentDefinitions": [],
    "kind": "Operation",
    "name": "UserListAllUsersQuery",
    "selections": (v0/*: any*/)
  },
  "params": {
    "cacheID": "215a01e0e40407c6b415fd489fdced5a",
    "id": null,
    "metadata": {},
    "name": "UserListAllUsersQuery",
    "operationKind": "query",
    "text": "query UserListAllUsersQuery {\n  users {\n    edges {\n      node {\n        id\n        username\n      }\n    }\n  }\n}\n"
  }
};
})();
(node as any).hash = 'ba3a3739b70406ee082ba351fd97ffe8';
export default node;
