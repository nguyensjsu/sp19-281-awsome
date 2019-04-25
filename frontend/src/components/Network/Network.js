import React,{Component} from 'react';
import Header from "../Common/Header";
import {Link} from "react-router-dom";


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
            user_profile_image :""
        }
    }

    render()
    {
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

                                <div className="col-lg-9 col-md-8 no-pd">
                                    My Network
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