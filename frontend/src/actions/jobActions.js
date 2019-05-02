import axios from 'axios';
import { POSTJOB, GET_ERRORS } from './types';
import {JOB_API} from "../components/constants/constants";
//Update Profile
export const postJob = (values, history) => dispatch => {
    axios
      .post(`${JOB_API}/recruiter/postJob`, values)
      .then(res => history.push('/home')
      )
      .catch(err =>
        dispatch({
          type: GET_ERRORS,
          payload: err.response.data
        })
      );
  };