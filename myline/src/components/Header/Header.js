import React, { useEffect, useState } from 'react'
import styled from 'styled-components';
import { HEADER_HEIGHT } from '../configs';
import liff from '@line/liff';
import {
  API_URL,
  setLineUserIdToken
} from '../../utils/ApiClient'

export const NavContainer = styled.div`
    display: flex;
    width: 100%;
    height: ${HEADER_HEIGHT};
    position: fixed;
    top: 0;
    z-index: 900;
    background-color: green;
    justify-content: space-between;
    flex-direction:column ;
    align-items: center;
    touch-action: none;
`;
export const HeaderContainer = styled.div`
    width: 100%;
    height:100% ;
    color: white;
    display: flex;
    padding:5px ;
    justify-content:space-between ;
    align-items: center;
    margin-left: 20px;
    margin-right: 20px;
`;

export const ButtonLogout = styled.div`
  width: 70px;
  height: 30px;
  border-radius:5px ;
  background-color:yellow ;
  justify-content:center ;
  align-items:center ;
  display:flex ;
  span{
    color:gray ;
    cursor: pointer;
  }

`;
const Header = () => {

  const [status, setStatus] = useState("Logout")

  useEffect(() => {

  }, [])

  const onHandler = () => {
    liff.init({
      liffId: API_URL.LIFF_ID,
    }).then(async () => {
      console.log("isLoggedIn : ", liff.isLoggedIn());
      if (liff.isLoggedIn()) {
        console.log("logout event")
        setStatus("Login")
        liff.logout()
      } else {
        console.log("login event")
        liff.login();
        setStatus("Logout")
      }
    }).catch((err) => {
      console.log(
        "%cERROR HEADER",
        "color: red",
        err.code,
        err.message
      );
    });
  }
  return (
    <>
      <NavContainer>
        <HeaderContainer>
          <h1>LINE TEST</h1>
          <ButtonLogout onClick={()=>onHandler()}><span>{status}</span></ButtonLogout>
        </HeaderContainer>
      </NavContainer>
    </>
  )
}

export default Header