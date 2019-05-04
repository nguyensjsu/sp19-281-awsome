import React, { Component } from "react";
import axios from "axios";
import { Redirect } from "react-router-dom";
import {JOB_API} from "../constants/constants";
import Header from "../Common/Header";
class JobDetails extends Component {
  constructor(props) {
    super(props);
    this.state = {
      profile: {}
    };
  }

  componentDidMount() {
    var splitPath = this.props.history.location.pathname.split("/");
    var id = splitPath[splitPath.length -1];
    axios
      .get(`${JOB_API}`+'/'+id)
      .then(response => {
        console.log("response.data" + JSON.stringify(response.data));
        this.setState({
          profile: response.data
          // profile: this.props.getProfile
        });
        console.log(this.state.profile);
      });
  }
  render() {
    // let redirectVar = null;
    // if (!localStorage.getItem("user_token")) {
    //   redirectVar = <Redirect to="/" />;
    // }

    return (
      <div>
        {/* {redirectVar} */}
        <Header />
        <div className="main-section pad-top-15">
                    <div className="container">
                        <div className="main-section-data">
                            <div className="row">
                                <div className="col-md-6 offset-md-2 pd-left-none no-pd">
                                    <div className="main-left-sidebar no-margin">
                                        <div className="user-data full-width">
                                            <div className="user-profile">
                                                
                                                <div className="user-specs">
                                                    <h3>{this.state.profile.title} At {this.state.profile.company} </h3>
                                                    <span>{this.state.profile.description}</span>
                                                </div>
                                            </div>
                                            <ul className="user-fw-status">
                                            <li>
                                                    <h4>Type of Job</h4>
                                                    <span>{this.state.profile.jobType}</span>
                                                </li>
                                                <li>
                                                    <h4>Industry</h4>
                                                    <span>{this.state.profile.industry}</span>
                                                </li>
                                                <li>
                                                    <h4>Function</h4>
                                                    <span>{this.state.profile.function}</span>
                                                </li>
                                                
                                            </ul>
                                        </div>
                                    </div>
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
      </div>
    );
  }
}
export default JobDetails;
