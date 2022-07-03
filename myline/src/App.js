import './App.css';
import { BrowserRouter as Router, Routes, Route, Outlet } from 'react-router-dom';
import { useState, useEffect } from 'react';
import { useDispatch, useSelector } from 'react-redux';
import Header from './components/Header/Header'
import MainLayout from './components/Layout/MainLayout'

import MainPage from './pages/Main/MainPage'

import UseLineLiff from './providers/UseLineLiff'
import {
  setLineToken
} from './utils/ApiClient'

function App() {

  const [loading, setLoading] = useState(true)
  const lineToken = UseLineLiff()

  useEffect(() => {
    console.log("app token = ", lineToken)
    if (lineToken.idtoken != "") {
      setLineToken(lineToken.idtoken,lineToken.accessToken)
      setLoading(false)
    }
  }, [lineToken])

  return (
    <div className="App">
      {loading ? <div>Loading...</div> :
        <Router>
          <Header />
          <MainLayout>
            <Routes>
              <Route path="/" element={<MainPage />} />
            </Routes>
          </MainLayout>
        </Router>
      }

    </div>
  );
}

export default App;
