import { Client } from "@modelcontextprotocol/sdk/client/index.js";
import { SSEClientTransport } from "@modelcontextprotocol/sdk/client/sse.js";
import { OAuthClientInformation, OAuthClientMetadata } from "@modelcontextprotocol/sdk/shared/auth.js";

const transport = new SSEClientTransport(new URL("http://localhost:8777"), {
    authProvider: {
        redirectUrl: "http://blah",
        clientMetadata: {} as unknown as OAuthClientMetadata,
        tokens() {
            return {
                access_token: "abcd:pj11017",
                refresh_token: "abcd",
                expires_in: undefined,
                token_type: "access_token",
                scope: "all"
            }
        },
        clientInformation() {
            return undefined
        },
        saveTokens(tokens) {},
        redirectToAuthorization() {},
        saveCodeVerifier() {},
        codeVerifier() {
            return "abc"
        },
    }
});

const client = new Client(
  {
    name: "example-client",
    version: "1.0.0"
  }
);

await client.connect(transport);

const tools = await client.listTools();
console.log(tools)

// Call a tool
const result = await client.callTool({
  name: "listTables",
  arguments: {
   
  }
});

console.log(result);

await client.close();