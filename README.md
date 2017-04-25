**Swift Note Taker and Searcher**

Qnote is a lightweight, easy-to-use tool for managing and finding quick notes.

**Key Features:**

*   Organize notes into collections (Books)
*   Create notes with your preferred editor (default: Vim)
*   Generate notes from web pages
*   Tag notes with hashtags
*   Edit and delete notes
*   Search with Bleve or Elasticsearch
*   Export notes to various formats (colored text, CSV, JSON)
*   Manage collections (delete, merge, split)
*   No configuration needed, just install and go
*   Compatible with Linux and macOS
*   Experimental CUI interface

**Getting Started:**

1.  **Install Go:** If you haven't already, install Go from [here](https://golang.org/).
2.  **Install Qnote:**  Use the following command: 

    ```bash
    go get github.com/anmil/quicknote/cmd/qnote 
    ```

**Managing Collections (Books):**

*   **Create a new collection:**

    ```bash
    qnote new book <collection name> 
    ```

*   **List all collections:**

    ```bash
    qnote ls books
    ```

*   **Delete a collection:**

    ```bash
    qnote rm book <collection name>
    ```

*   **Merge collections:**

    ```bash
    qnote merge <collection to delete> <collection to move notes to>
    ```

*   **Split collections:**

    ```bash
    qnote split query <collection name> <search query>
    qnote split ids <collection name> <note IDs...>
    ```

**Creating and Managing Notes:**

*   **Create a new note:**

    ```bash
    qnote new note 
    ```

*   **Create a note from a URL:**

    ```bash
    qnote new url <URL>
    ```

*   **List notes:**

    ```bash
    qnote ls notes 
    qnote ls notes all 
    ```

*   **Edit a note:**

    ```bash
    qnote edit note <note ID>
    ```

*   **Delete a note:**

    ```bash
    qnote rm note <note ID>
    ```

**Searching Notes:**

*   **Search using phrase prefix:**

    ```bash
    qnote search query 
    ```

*   **Search using advanced query syntax:**

    ```bash
    qnote search -q <search query> 
    ```

**Re-indexing:**

*   If you need to rebuild the search index, use:

    ```bash
    qnote search reindex
    ```


**Backup and Restore:**

*   To back up, copy the `qnote.db` file from the data directory (`$HOME/.config/quicknote` on Linux and `$HOME/Library/Application Support/quicknote` on macOS).
*   Exporting to CSV or JSON is also an option.


**Command Help:**

*   Get information on any command using the `help` command:

    ```bash
    qnote help <command>
    ```

**Example Usage:**

```
$ qnote help
$ qnote new book Work
$ qnote new note
$ qnote search -q "book:Work AND tag:projectx"
$ qnote ls notes all -f json
```