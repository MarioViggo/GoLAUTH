import React from 'react';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import SignIn from './pages/signIn';
import GlobalStyles from './styles/global';

import 'bootstrap/dist/css/bootstrap.min.css';

function App() {
  const logOut = () => {
    localStorage.clear();
    window.location.reload();
  }
  return (

    <div className="App">
      <Router>
        <GlobalStyles />
          <SignIn />
        <Routes>
          <Route path="/signin" element={<SignIn />} />
        </Routes>
      </Router>
    </div >
  );
}

export default App;
