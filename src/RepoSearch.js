import React, { Component } from "react"

class RepoSearch extends Component {
  constructor() {
    super()
    this.state = {
      files: "",
      repoURL: "",
      searchTerm: "",
      toggle: true,
      searchResults: ""
    }

    this.handleSubmit = this.handleSubmit.bind(this)
    this.handleChange = this.handleChange.bind(this)
    this.handleLookup = this.handleLookup.bind(this)
  }
  handleSubmit(event) {
    fetch("http://localhost:5000", {
      method: "POST",
      headers: {
        "Content-Type": "application/json"
      },
      body: JSON.stringify({ "url": this.state.repoURL })
    })
    .then(res => res.json())
    .then(files => {
      this.setState({ files })
    })
    this.setState({ toggle: false })
    this.refs.repoSearch.value = ""
    event.preventDefault()
  }
  handleLookup(event) {
    event.preventDefault()
    const map = new Map(Object.entries(this.state.files))
    const l = lookup(map, this.state.searchTerm)
    this.setState({ searchResults: l })
  }
  handleChange(event) {
    this.setState({ [event.target.name]: event.target.value })
  }
  render() {
    let resultsAreIn = false
    if (this.state.toggle === false && this.state.searchResults) {
      resultsAreIn = true
    }
    return (
      <div>
        {this.state.toggle
          ? <form onSubmit={this.handleSubmit}>
              <input type="text" placeholder="Repository url" ref="repoSearch" name="repoURL" onChange={this.handleChange}/>
            </form>
          : <form onSubmit={this.handleLookup}>
              <input type="text" placeholder="Search.." name="searchTerm" onChange={this.handleChange}/>
              {resultsAreIn &&
                <Results matches={this.state.searchResults}/>
              }
             </form>
        }
      </div>
    )
  }
}

const Results = (props) => {
  let l = []
  console.log(props.matches)
  for (const [key, val] of props.matches.entries()) {
    if (val.size > 0) {
      for (const [nk, nv] of val.entries()) {
        l.push(key + ":" + nk + " " + nv)
      }
    }
  }
  const results = l.map((result) =>
    <li key={result}>{result}</li>
  )
  return (
    <div>
      <p>Results are in!</p>
      <ul>{results}</ul>
    </div>
  )
}

function lookup(files, term) {
  const l = new Map()
  for (const [key, val] of files.entries()) {
    let map = searchText(val, term)
    l.set(key, map)
  }
  return l
}

function searchText(text, term) {
  let lines = text.split("\n")
  const map = new Map()
  for (let i = 0; i < lines.length; i++) {
    if (lines[i].includes(term)) {
      map.set(i, lines[i])
    }
  }
  return map
}

export default RepoSearch
