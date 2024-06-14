import { useState, useEffect } from 'react';
import axios from "axios";
import {useCourses, CourseDetail, CourseList} from './Cursos';

const Home = ({ token }) => {
  
    const [courses, searchTerms, setSearchTerms] = useCourses(token);
  
    const [detailCourse, setDetailCourse] = useState(null);
    
    if (detailCourse) {
      return <CourseDetail course={detailCourse} hideDetail={() => setDetailCourse(null)} />;
    }
  
    return <CourseList courses={courses} setDetailCourse={setDetailCourse} searchTerms={searchTerms} setSearchTerms={setSearchTerms} />;
  
  };
  
  export default Home;