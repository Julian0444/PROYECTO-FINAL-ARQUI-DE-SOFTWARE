import React, { useState, useEffect } from 'react';
import axios from 'axios';
import { BrowserRouter as Router, Route, Routes, Link} from "react-router-dom";
import Home from "./Home"
import Login from "./Login"

function App(){
  
  const [token, setToken] = useState(null)
  return (
      <Router>
        <Routes>
          <Route exact path = "/" element={<Home token={token}/>}/>
          <Route path= "/login" element={<Login token={token} setToken={setToken}/>}/>
        </Routes>
      </Router>
    );
  }
  
  export default App;

