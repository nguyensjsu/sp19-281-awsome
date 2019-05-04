import React from 'react';
import {Route,BrowserRouter} from 'react-router-dom';
import './App.css';
import HomePage from "./components/Home/Home";
import ApplicantHome from "./components/Home/ApplicantHome"
import RecruiterHome from "./components/Home/RecruiterHome"
import Network from "./components/Network/Network";
import Profile from "./components/profile/profile";
import PostJob from "./components/UserJobs/postJob";
import ListJobs from './components/UserJobs/listJobs';
import JobDetails from './components/UserJobs/jobDetails';

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
                    <Route path="/postjobs" component={PostJob}/>
                    <Route path="/listJobs" component={ListJobs}/>
                    <Route path="/jobDetails/:id" component={JobDetails}/>
                </div>
          </div>
      </BrowserRouter>
  );
}

export default App;
