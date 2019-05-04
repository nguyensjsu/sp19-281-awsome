import React, { Component } from 'react';
import Header from "../Common/Header";
import "../../css/JobSearch.css"
import axios from 'axios'
import {JOB_API} from "../constants/constants";
var jwtDecode = require('jwt-decode');


class ListJobs extends Component {
    constructor(props) {
        super(props);
        this.state = {
            ListJobs: [],
        }
        this.onUnsave = this.onUnsave.bind(this);
    }

    async componentDidMount() {
        console.log("in component did mount save jobs")
        // var decoded = jwtDecode(localStorage.getItem("userToken"));
        // const applicantId = decoded.id
        const response = await axios.get(`${JOB_API}`)
        console.log(response.data)
        this.setState({
            ListJobs: response.data
        })

    }
    async onUnsave(id, e) {
        console.log("in unsave")
        var decoded = jwtDecode(localStorage.getItem("userToken"));
        const applicantId = decoded.id;
        const jobId = id
        const response = await axios.delete(`${JOB_API}`, { params: { applicantId: applicantId, jobId: jobId } })
        if (response.status === 200) {
            window.location.reload()
        }
    }
    render() {
        var details = null
        console.log(this.state.ListJobs)

        if (this.state.ListJobs !== "") {
            console.log("something")
            var details = this.state.ListJobs.map((value, i) => {
                return (
                    <div>
                        {/* <h4><strong>{value.title}</strong> */}
                        <h4><strong><a href={'https://cmpe281-awsome.herokuapp.com/jobDetails/' + value.id} title={value.headline}>{value.title}</a></strong>
                            {/* <button tabindex="1" onClick={this.onUnsave.bind(this, value.jobId)} className="job-page-job-unsave"><span class="glyphicon glyphicon-bookmark"></span></button> */}
                            <button tabindex="1" onClick={this.onUnsave.bind(this, value.title)} className="job-page-job-unsave"><span class="glyphicon glyphicon-bookmark"></span></button>
                        </h4>
                        <h5>{value.company}</h5>
                        <h6>{value.industry},{value.jobType}</h6>
                        {/* <h6><font color="gray">Posted on: {value.dateSaved.slice(0, 10)}</font></h6> */}
                        <h6><font color="gray">Description: {value.description}</font></h6>
                        <hr></hr>
                    </div>
                )

            })

        }
        else {
            console.log("nothing")
        }


        return (
            <div>
                <Header/>
                <div className="containerListJobs">

                    <br></br>
                    <br></br>
                    <br></br>
                    <br></br>
                    <div className="container">
                        <div className="row" >
                            <div className="col-md-2" >

                            </div>
                            <div className="col-md-2" ></div>
                            <div className="col-md-2" ></div>
                            <div className="col-md-2" ></div>
                        </div>
                        <div className="row">
                            <div className="col-md-1" >
                                <div className="card searchresultscards1">
                                    <div className="card-body">
                                        <h4 className="fontsHeading">All Jobs</h4>
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
        )
    }
}
export default ListJobs
