import React, { Component } from "react";
import { Button, FormGroup, FormControl, FormLabel } from "react-bootstrap";
import "./Login.css";
import axios from 'axios';
const axiosGraphQL = axios.create({
    baseURL: 'http://localhost:7777/query',
  });

export default class Login extends Component {
  constructor(props) {
    super(props)
    

    this.state = {
      email: "",
      password: ""
    };
    this.error = <span></span>
  }

  validateForm() {
    return this.state.email.length > 0 && this.state.password.length > 0;
  }

  handleChange = event => {
    this.setState({
      [event.target.id]: event.target.value
    });
  }
  login() {
    
    axiosGraphQL.post('', {query: `
      query{
        users{
          user_id
          email
          password
        }
        }

    `}).then(result => {
      let email = this.state.email
      let password = this.state.password
      if(result.data.data.users.some(user=>user.email === email && user.password ===password)){
        window.location.href = '/'
      } else {
        document.querySelector('span').textContent="Wrong email or password!"
      }
    })
  }

  render() {
    return (
      <div className="Login">
          <FormGroup controlId="email" bsSize="large">
            <FormLabel>Email</FormLabel><br></br>
            <FormControl
              autoFocus
              type="email"
              value={this.state.email}
              onChange={this.handleChange}
            />
          </FormGroup><br></br>
          <FormGroup controlId="password" bsSize="large">
            <FormLabel>Password</FormLabel><br></br>
            <FormControl
              value={this.state.password}
              onChange={this.handleChange}
              type="password"
            />
          </FormGroup><br></br>
          <Button
            block
            bsSize="large"
            disabled={!this.validateForm()}
            type="submit"
            onClick={() => this.login()}
          >
            Login
          </Button><br></br>
          {this.error}
      </div>
    );
  }
}