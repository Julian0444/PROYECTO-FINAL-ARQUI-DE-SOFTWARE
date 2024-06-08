import React, { useState, useEffect } from 'react';
import axios from 'axios';

const useCourses = (token) => {
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
        console.error('Error al obtener cursos:', error);
      });
  }, [searchTerms, token]);

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


const LoginPage = () => {
  return (
    <div>
      <h1>Inicio de sesion</h1>
      <form>
        <label>
          Email:
          <input type="text" name="email" />
        </label>
        <label>
          Contraseña:
          <input type="password" name="password" />
        </label>
        <button type="submit">Iniciar sesion</button>
      </form>
    </div>
  );

}


const CreateUserPage = ({ setToken }) => {
  const [name, setName] = useState('');
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const handleSubmit = async (event) => {
    event.preventDefault();

    const targetUrl = 'http://localhost:8080/auth/register';
    axios.post(targetUrl, {
      name,
      email,
      password,
    }, {
      headers: {
        'Content-Type': 'application/json',
      },
    }).then(response => {

      if (response.data.data.token) {
        setToken(response.data.data.token);
      }

    })

  };

  return (
    <div>
      <h1>Crear usuario</h1>
      <form onSubmit={handleSubmit}>
        <label>
          Nombre:
          <input type="text" name="name" value={name} onChange={(e) => setName(e.target.value)} />
        </label>
        <label>
          Email:
          <input type="text" name="email" value={email} onChange={(e) => setEmail(e.target.value)} />
        </label>
        <label>
          Contraseña:
          <input type="password" name="password" value={password} onChange={(e) => setPassword(e.target.value)} />
        </label>
        <button type="submit">Crear usuario</button>
      </form>
    </div>
  );
}


const App = () => {
  const [token, setToken] = useState(null)

  const [courses, searchTerms, setSearchTerms] = useCourses(token);

  const [detailCourse, setDetailCourse] = useState(null);

  if (token === null) {
    return <CreateUserPage setToken={setToken} />;
  }

  if (detailCourse) {
    return <CourseDetail course={detailCourse} hideDetail={() => setDetailCourse(null)} />;
  }

  return <CourseList courses={courses} setDetailCourse={setDetailCourse} searchTerms={searchTerms} setSearchTerms={setSearchTerms} />;

};

export default App;