import React, { useState, useEffect } from 'react';
import axios from 'axios';

const useCourses = () => {
  const [courses, setCourses] = useState([]);
  const [searchTerms, setSearchTerms] = useState("");
  useEffect(() => {
    const targetUrl = 'http://localhost:8080/courses' + (searchTerms ? `?search=${searchTerms}` : '');
    axios.get(targetUrl)
      .then(response => {
        if (response.data.Data !== null) {
          setCourses(response.data.Data);
        }
      })
      .catch(error => {
        console.error('Error al obtener cursos:', error);
      });
  }, [searchTerms]);

  return [courses, searchTerms, setSearchTerms];
}


const CourseDetail = ({ course, hideDetail }) => {
  return (
    <div>
      <h1>{course.CourseName}</h1>
      <p>{course.CourseDescription}</p>
      <p>Precio: ${course.CoursePrice}</p>
      <p>Duracion: {course.CourseDuration} days</p>
      <p>Fecha de inicio: {course.CourseInitDate}</p>
      <p>Capacidad: {course.CourseCapacity}</p>
      <img src={course.CourseImage} alt={course.CourseName} />
      <h2>{course.CourseName}</h2>
      <p>{course.CourseDescription}</p>
      <button onClick={hideDetail}>Cerrar</button>
    </div>
  );
};

const CourseList = ({ courses, setDetailCourse, searchTerms, setSearchTerms }) => {

  const handleChange = (event) => {
    setSearchTerms(event.target.value); // Update state when input changes
  };
  return (
    <div>
      <input type="text" value={searchTerms} onChange={handleChange} />;
      <h2>Listado de Cursos</h2>
      <ul>
        {courses.map(curso => (
          <li key={curso.Id}>
            <h3>{curso.CourseName}</h3>
            <p>{curso.CourseDescription}</p>
            <button onClick={() => setDetailCourse(curso)}>Ver detalle</button>
          </li>
        ))}
      </ul>
    </div>
  );
}


const CourseCatalog = () => {
  const [courses, searchTerms, setSearchTerms] = useCourses();

  const [detailCourse, setDetailCourse] = useState(null);

  if (detailCourse) {
    return <CourseDetail course={detailCourse} hideDetail={() => setDetailCourse(null)} />;
  }

  return <CourseList courses={courses} setDetailCourse={setDetailCourse} searchTerms={searchTerms} setSearchTerms={setSearchTerms} />;

};

export default CourseCatalog;