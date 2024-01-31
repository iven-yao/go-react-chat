import React, { useEffect, useState } from 'react';
import logo from './logo.svg';
import './App.css';
import axios from 'axios';
import User from './types/User';
import UserServices from './services/UserService';



function App() {
  const [users, setUsers] = useState<User[]>([]);
  useEffect(() => {
    // UserServices.getAll().then(res => setUsers(res.data)).catch(err => console.error(err))
    UserServices.get(6).then(res => setUsers([res.data])).catch(err => console.error(err))
  },[])

  return (
    <div className="App">
      <header className="App-header">
        <img src={logo} className="App-logo" alt="logo" />
        <p>
          Edit <code>src/App.tsx</code> and save to reload.
        </p>
        
        <div>{users.map(user => <p key={user.ID.toString()}>{user.username}</p>)}</div>
      </header>
    </div>
  );
}

export default App;
