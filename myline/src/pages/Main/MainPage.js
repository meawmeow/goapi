import React from 'react'
import styled from 'styled-components';
import Wrapper from '../../components/Layout/Wrapper'

import {
  LineUserIdToken,
  LineUserAccessToken
} from '../../utils/ApiClient'
export const Container = styled.div`
  width: 100%;
  height: auto;
  background-color:pink ;
  word-wrap: break-word;
`;


const MainPage = () => {
  return (
    <Wrapper>
      <Container>
        <h3>
          ID TOKEN
        </h3>
        <span>
          {LineUserIdToken}
        </span>
        <hr/>
        <h3>
          ID ACCESS TOKEN
        </h3>
        <span>
          {LineUserAccessToken}
        </span>
  

      </Container>
    </Wrapper>
  )
}

export default MainPage