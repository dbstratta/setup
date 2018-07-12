""" Plugins
call plug#begin()

" Code consistency
Plug 'editorconfig/editorconfig-vim'

" Syntax and language integration
Plug 'sheerun/vim-polyglot'

call plug#end()

""" Core
set encoding=utf-8
set number

""" Keybindings
map <Space> <leader>

" Remap saving and quitting
nnoremap <leader>w :w<CR>
nnoremap <leader>q :q<CR>
