import React, { useEffect, useState } from 'react';
import axios from "axios";
import { useNavigate } from "react-router-dom";
import './Cursos.css';


export const useCourses = (token) => {
  const navigate = useNavigate();
  const [courses, setCourses] = useState([]);
  const [searchTerms, setSearchTerms] = useState("");
  useEffect(() => {
    const targetUrl = 'http://localhost:8080/courses' + (searchTerms ? `?search=${searchTerms}` : '');
    axios.get(targetUrl, {
      headers: {
        'Authorization': `Bearer ${token}`,
      },
    })
      .then(response => {
        if (response.data.Data !== null) {
          setCourses(response.data.Data);
        }
      })
      .catch(error => {
        if (error.response.status === 401) {
          navigate("/login");
        }
        console.error('Error al obtener cursos:', error);
      });
  }, [searchTerms, token]);

  return [courses, searchTerms, setSearchTerms];
}


export const CourseDetail = ({ course, hideDetail }) => {
  return (
    <div className="course-detail container">
      <h1>{course.CourseName}</h1>
      <p>{course.CourseDescription}</p>
      <p>Precio: ${course.CoursePrice}</p>
      <p>Duracion: {course.CourseDuration} horas</p>
      <p>Fecha de inicio: {course.CourseInitDate}</p>
      <p>Capacidad: {course.CourseCapacity}</p>
      <img src={course.CourseImage} alt={course.CourseName} />
      <h2>{course.CourseName}</h2>
      <p>{course.CourseDescription}</p>
      <button onClick={hideDetail}>Cerrar</button>
    </div>
  );
};

export const CourseList = ({ courses, setDetailCourse, searchTerms, setSearchTerms }) => {

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
