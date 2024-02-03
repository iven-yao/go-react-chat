import React from 'react';
import './App.css';
import Login from './components/Login';
import { BrowserRouter, Route, Routes } from 'react-router-dom';
import ProtectedRoute, { ProtectedRouteProps } from './ProtectedRoute';
import Main from './components/Main';
import Cookies from "universal-cookie";
const cookies = new Cookies();


const defaultProtectedRouteProps: Omit<ProtectedRouteProps, 'outlet'> = {
  isAuth: cookies.get("token") !== undefined,
  redirectPath: '/login'
}

function App() {
  return (
    <div>
      <BrowserRouter>
        <Routes>
          <Route path="/login" element={<Login />}/>
          <Route path="/" element={<ProtectedRoute {...defaultProtectedRouteProps} outlet={<Main />}/>} />
        </Routes>
      </BrowserRouter>
    </div>
    
  );
}

export default App;
