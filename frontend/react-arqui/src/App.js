import React, { useState, useEffect } from 'react';
import axios from 'axios';

const CursosList = () => {
  const [cursos, setCursos] = useState([]);

  useEffect(() => {
    axios.get('http://localhost:8080/courses')
      .then(response => {
        setCursos(response.data.Data);
      })
      .catch(error => {
        console.error('Error al obtener cursos:', error);
      });
  }, []);

  return (
    <div>
      <h2>Listado de Cursos</h2>
      <ul>
        {cursos.map(curso => (
          <li key={curso.Id}>
            <h3>{curso.CourseName}</h3>
            <p>{curso.CourseDescription}</p>
          </li>
        ))}
      </ul>
    </div>
  );
};

export default CursosList;