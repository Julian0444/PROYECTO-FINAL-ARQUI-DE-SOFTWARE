import { useState, useEffect } from 'react';
import axios from "axios";
import {Cursos, CourseDetail, CourseList} from './Cursos';

const Home = () => {
    const [token, setToken] = useState(null)
  
    const [courses, searchTerms, setSearchTerms] = Cursos(token);
  
    const [detailCourse, setDetailCourse] = useState(null);
  
    if (token === null) {
      return <CreateUserPage setToken={setToken} />;
    }
  
    if (detailCourse) {
      return <CourseDetail course={detailCourse} hideDetail={() => setDetailCourse(null)} />;
    }
  
    return <CourseList courses={courses} setDetailCourse={setDetailCourse} searchTerms={searchTerms} setSearchTerms={setSearchTerms} />;
  
  };
  
  export default Home;