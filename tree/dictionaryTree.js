// 字典树
var Trie = function () {
    this.child = {}
}

Trie.prototype.insert = function(word) {
    let node = this.child
    for (const letter of word) {
        if (!node[letter]) {
            node[letter] = {}
        }
        node = node[letter]
    }
    node.isWord = true
}

Trie.prototype.find = function(word) {
    const node = this.findPrefix(word)
    return node !== undefined && node.isWord !== undefined
}

Trie.prototype.findPrefix = function(prefix) {
    let node = this.child
    for (const letter of prefix) {
        if (!node[letter]) {
            return false
        }
        node = node[letter]
    }
    return node
}