import React from 'react';
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import Paper from '@material-ui/core/Paper';
import details from '../details.svg'
import { Link } from "react-router-dom";
import { Component } from "react";
import axios from 'axios';

const axiosGraphQL = axios.create({
  baseURL: 'http://localhost:7777/query',
});

const GET_BOOKS = `
query{
    books {
      book_id
      title
      author_id
      owner_id
      location_id
      is_free
    }
  }
`;

class BookTable extends Component {
  authors = new Map();
  users = new Map();
  locations = new Map();
  state = {
    json: {},
    loading: 0,
};
    componentDidMount() {
        this.onFetch();
      }
    onFetch() {
        axiosGraphQL
        .post('', { query: GET_BOOKS })
        .then(result =>  {
          this.setState({ json: result, loading: 0 })
          return result;
        })
        .then(result => {
          Promise.all([...new Set(result.data.data.books.map(book => book.author_id))].map(author => axiosGraphQL.post('', {
            query: `
            query{
              author(author_id: ${author}){
                author_id
                name
                surname
              }
            }
            `
          }))).then(arr => {
            arr.forEach(e => {
              this.authors.set(e.data.data.author.author_id, `${e.data.data.author.name} ${e.data.data.author.surname}`)
            })
            return true;
          })
          .then(_ => this.setState({ loading: this.state.loading + 1 }));
          
          return result;
        })
        .then(result => {
          Promise.all([...new Set(result.data.data.books.map(book => book.owner_id))].map(owner => axiosGraphQL.post('', {
            query: `
            query{
              user(user_id: ${owner}){
                user_id
                name
                surname
              }
            }
            `
          }))).then(arr => {
            arr.forEach(e => {
              this.users.set(e.data.data.user.user_id, `${e.data.data.user.name} ${e.data.data.user.surname}`)
            })
            return true;
          })
          .then(_ => this.setState({ loading: this.state.loading + 1 }));
          
          return result;
        })
        .then(result => {
          Promise.all([...new Set(result.data.data.books.map(book => book.location_id))].map(location => axiosGraphQL.post('', {
            query: `
            query{
              location(location_id: ${location}) {
                location_id
                building
                room
              }
            }
            `
          }))).then(arr => {
            arr.forEach(e => {
              this.locations.set(e.data.data.location.location_id, `${e.data.data.location.building} ${e.data.data.location.room}`)
            })
            return true;
          })
          .then(_ => this.setState({ loading: this.state.loading + 1 }));
          
          return result;
        });
    }
    render() {
      const { loading, json } = this.state;
      if (loading < 3) {
          return <span>Loading...</span>
      }
  return (
    <Paper className="paper">
      <Table className="BookTable">
        <TableHead>
          <TableRow>
            <TableCell>Tytuł</TableCell>
            <TableCell align="center">Autor</TableCell>
            <TableCell align="center">Właściciel</TableCell>
            <TableCell align="center">Lokalizacja</TableCell>
            <TableCell align="center">Dostępność</TableCell>
            <TableCell align="center">Szczegóły</TableCell>
          </TableRow>
        </TableHead>
        <TableBody>
          {json.data.data.books.map(book => (
            <TableRow key={book.book_id}>
              <TableCell scope="row">
              {book.title}
              </TableCell>
              <TableCell align="center">{this.authors.get(book.author_id)}</TableCell>
              <TableCell align="center">{this.users.get(book.owner_id)}</TableCell>
              <TableCell align="center">{this.locations.get(book.location_id)}</TableCell>
              <TableCell align="center">{book.is_free ? "Free": "Not available"}</TableCell>
              <TableCell align="center">
                <Link to={`/book/${book.book_id}`}>
                    <img src={details} className="img" alt="loupe"/>
                </Link>
              </TableCell>
            </TableRow>
          ))}
        </TableBody>
      </Table>
    </Paper>
  );}
}

export default BookTable;