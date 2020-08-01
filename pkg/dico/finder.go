package dico

func (d *Dico) Find(word string) bool {
	if _, ok := d.words[word]; ok {
		return true
	}
	return false
}
