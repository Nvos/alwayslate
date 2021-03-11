// your-app-name/src/RelayEnvironment.js
import {Environment, Network, RecordSource, Store} from 'relay-runtime';

function fetchQuery(operation: any, variables: any, cacheConfig: any, uploadables: any) {
    // @ts-ignore
    return fetch("http://localhost:8080/query", {
        method: 'POST',
        headers: {
            Accept: 'application/json',
            'Content-Type': 'application/json',
        }, // Add authentication and other headers here
        body: JSON.stringify({
            query: operation.text,
            variables,
        }),
    }).then((response) => response.json())
}

// Export a singleton instance of Relay Environment configured with our network function:
export default new Environment({
    network: Network.create(fetchQuery),
    store: new Store(new RecordSource()),
});