# subt

Helps to rename srt files in a directory. Computes levenshtein distance between each srt and other file types to find the best match.

Without arguments, it suggests the 3 best matches, user needs to choose for each srt file.

With `-show` argument, it only prints the srt files and their first best match.

With `-okFirst` argument, it renames each srt files to their first best match.
