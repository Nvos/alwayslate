import React, { useState, Suspense } from 'react'
import {
  RelayEnvironmentProvider,
} from 'react-relay/hooks';
import './App.css'
import client from "./graphql/client";
import UserList from "./UserList/UserList";

const App = () => {
  return (
    <RelayEnvironmentProvider environment={client}>
      <Suspense fallback={'Loading...'}>
        <UserList />
          Test
      </Suspense>
    </RelayEnvironmentProvider>
  )
}

export default App
