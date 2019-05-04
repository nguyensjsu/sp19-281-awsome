import React,{Component} from 'react';
import Header from "../Common/Header";
import {Link} from "react-router-dom";
import "../../css/JobSearch.css"
import {JOB_API, USER_API} from "../constants/constants";
import axios from "axios";
var jwtDecode = require('jwt-decode');


class Network extends Component
{
    constructor(props)
    {
        super(props);
        this.state={
            connections:[],
            totalConnections:0,
            fname:"",
            lname:"",
            headline : "",
            user_profile_image :"",
            users:[]
        }
        this.onUnsave = this.onUnsave.bind(this);
    }

    async componentDidMount() {
        console.log("in component did mount save jobs")
        // var decoded = jwtDecode(localStorage.getItem("userToken"));
        // const applicantId = decoded.id
        const response = await axios.get(`${USER_API}`)
        console.log(response.data)
        this.setState({
            users: response.data
        })

    }


    async onUnsave(id, e) {
        console.log("in unsave")
        var decoded = jwtDecode(localStorage.getItem("userToken"));
        const email = decoded.email;
        const jobId = id
        // const response = await axios.delete(`${JOB_API}`, { params: { applicantId: applicantId, jobId: jobId } })
        // if (response.status === 200) {
        //     window.location.reload()
        // }
    }

    render()
    {

        var details = null
        console.log(this.state.users)

        if (this.state.users !== "") {
            console.log("something")
            var details = this.state.users.map((value, i) => {
                return (
                    <div>
                        {/* <h4><strong>{value.title}</strong> */}
                        <h4><strong><a href="#">{value.firstname}</a></strong>
                            {/* <button tabindex="1" onClick={this.onUnsave.bind(this, value.jobId)} className="job-page-job-unsave"><span class="glyphicon glyphicon-bookmark"></span></button> */}
                            <button tabindex="1" onClick={this.onUnsave.bind(this, value.title)} className="job-page-job-unsave"><span class="glyphicon glyphicon-bookmark"></span></button>
                        </h4>
                        <h5>{value.lastname}</h5>
                        <h6>{value.email},{value.company}</h6>
                        {/* <h6><font color="gray">Posted on: {value.dateSaved.slice(0, 10)}</font></h6> */}
                        <h6><font color="gray">About Me: {value.aboutme}</font></h6>
                        <hr></hr>
                    </div>
                )

            })

        }
        else {
            console.log("nothing")
        }

        return(
            <div>
                <Header/>

                <div className="main-section pad-top-15">
                    <div className="container">
                        <div className="main-section-data">
                            <div className="row">
                                <div className="col-lg-3 col-md-4 pd-left-none no-pd">
                                    <div className="main-left-sidebar no-margin">
                                        <div className="user-data full-width">
                                            <div className="user-profile">
                                                <div className="username-dt">
                                                    <div className="usr-pic">
                                                        <img src={this.state.user_profile_image} alt=""/>

                                                    </div>
                                                </div>
                                                <div className="user-specs">
                                                    <h3>{this.state.fname} {this.state.lname} </h3>
                                                    <span>{this.state.headline}</span>
                                                </div>
                                            </div>
                                            <ul className="user-fw-status">
                                                <li>
                                                    <h4>Connections</h4>
                                                    <span>{this.state.totalConnections}</span>
                                                </li>
                                                <li>
                                                    <Link to="/profile">View Profile</Link>
                                                </li>
                                            </ul>
                                        </div>
                                    </div>
                                </div>

                                <div className="containerListJobs">

                                    <br></br>
                                    <br></br>
                                    <br></br>
                                    <br></br>
                                    <div className="container">
                                        <div className="row" >
                                            <div className="col-md-8" >

                                            </div>
                                            <div className="col-md-2" ></div>
                                            <div className="col-md-2" ></div>
                                            <div className="col-md-2" ></div>
                                        </div>
                                        <div className="row">
                                            <div className="col-md-8" >
                                                <div className="card searchresultscards1">
                                                    <div className="card-body">
                                                        <h4 className="fontsHeading">All Users</h4>
                                                        <hr></hr>
                                                        {details}
                                                    </div>
                                                </div>
                                            </div>
                                            <div className="col-md-1" ></div>
                                            <div className="col-md-1" ></div>
                                            <div className="col-md-1" ></div>
                                        </div>
                                    </div>
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default Network;