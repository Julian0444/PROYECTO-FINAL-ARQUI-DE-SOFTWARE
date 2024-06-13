import React, { useState, useEffect } from 'react';
import axios from 'axios';
import { BrowserRouter as Router, Route, Routes, Link} from "react-router-dom";
import Home from "./Home"
import Curso from "./Cursos"
import Login from "./Login"

function App(){
  return (
      <Router>
        <Routes>
          <Route exact path = "/" element={<Home/>}/>
          <Route path= "/cursos" element={<Curso/>}/>
          <Route path= "/login" element={<Login/>}/>
        </Routes>
      </Router>
    );
  }
  
  export default App;

