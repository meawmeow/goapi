import React from 'react'
import styled from 'styled-components';

export const Container = styled.div`
  width: 100%;
  padding:10px ;
`;
const Wrapper = ({ children }) => {
    return (
        <Container>{children}</Container>
    )
}

export default Wrapper