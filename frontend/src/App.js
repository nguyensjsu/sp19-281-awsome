import React from 'react';
import {Route,BrowserRouter} from 'react-router-dom';
import './App.css';
import HomePage from "./components/Home/Home";
import ApplicantHome from "./components/Home/ApplicantHome"
import RecruiterHome from "./components/Home/RecruiterHome"
import Network from "./components/Network/Network";
import Profile from "./components/Profile/profile";
import PostJob from "./components/UserJobs/postJob";

function App() {
  return (
      <BrowserRouter>
          <div className="App">
                <div>
                    {/*Render Different Component based on Route*/}
                    <Route exact path="/" component={HomePage}/>
                    <Route path="/applicanthome" component={ApplicantHome} />
                    <Route path="/recruiterhome" component={RecruiterHome} />
                    <Route path="/mynetwork" component={Network} />
                    <Route path="/profile" component={Profile}/>
                    <Route path="/jobshome" component={PostJob}/>
                </div>
          </div>
      </BrowserRouter>
  );
}

export default App;
