import React,{Component} from 'react';
import Header from "../Common/Header";
import {Link} from "react-router-dom";


class UserJobs extends Component
{
    constructor(props)
    {
        super(props);
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

                                <div className="col-lg-9 col-md-8 no-pd">
                                    UserJobs
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default UserJobs;