import React, {FC, Suspense} from 'react';
// import { graphql } from 'react-relay';
import graphql from 'babel-plugin-relay/macro';
import {useLazyLoadQuery} from "react-relay/hooks";
import {UserListAllUsersQuery} from '../__generated__/UserListAllUsersQuery.graphql';
import BrokenDependency, {ReactSwitchProps} from "react-switch";

const Switch: FC<ReactSwitchProps> = (BrokenDependency as any).default;

const Query = graphql`
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
`

const UserList = () => {
    const data = useLazyLoadQuery<UserListAllUsersQuery>(Query, {})

    return <div>
        {data.users?.edges?.map(user => <div>
            {user!.node.id}
            {user!.node.username}
        </div>)}
        <Switch checked={true} onChange={() => {}} />
    </div>
}

export default UserList;