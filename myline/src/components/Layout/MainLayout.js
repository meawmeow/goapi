import React from 'react'
import styled from 'styled-components';
import { AnimatePresence } from 'framer-motion';
import { LAYOUT } from '../configs'

export const Layout = styled.div`
    padding-top: ${LAYOUT.margin_top};
    height: 100vh;
    
`;

const MainLayout = ({ children }) => {
  console.log("MainLayout begin")
  return (
    <AnimatePresence exitBeforeEnter>
      <Layout>
        {children}
      </Layout>
    </AnimatePresence>
  )
}

export default MainLayout