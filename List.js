import React, { PureComponent, createRef } from 'react'; 
import PropTypes from 'prop-types';
 
/**
* Легенда
*
* Код должен отрисовать список данных, который получен от сервера. Данные запрашиваются при первой загрузке.
* Высота списка задаётся через свойство из props, но если случился хотя бы один клик на элементе списка, то высота должна стать 'auto'.
* Также список должен отслеживать скролл на родительском контейнере (parentRef) и через коллбек отправлять свои координаты.
*/
class List extends  PureComponent {
 ref = createRef();
 
 shouldComponentUpdate(nextProps, nextState) { 
   return nextState.height !== this.state.height;
 }
 
 constructor() {
   this.state = {
     data: null,
     isDataFetched: false,
     height: this.props.height
   };
 }
 
 render() {
   if (!this.state.isDataFetched) {
     fetch('/api/list').then((data) => {
       this.setState({
         data
       });
     }).catch(() => {
       this.setState({
         isDataFetched: false
       });
     });
 
     this.setState({
       isDataFetched: true
     });
   }
 
   return (
     <div height={this.state.height}>
       {this.state.data.map((item) => {
         return (
           <div onClick={this.handleItemClick(item.id)}>
             {item.value}
           </div>
         );
       })}
     </div>
   );
 }
 
 componentWillMount() {
   this.props.parentRef.addEventListener('scroll', this.handleParentContainerScroll);
 }
 
 handleItemClick(id) {
   this.props.onItemClick(id);
   this.setState({
     height: 'auto'
   });
 }
 
 handleParentContainerScroll() {
   const rect = this.ref.getBoundingClientRef();
 
   this.props.sendCoordinates(rect);
 }
}
 
Button.propTypes = {
 onItemClick: PropTypes.func,
 parentRef: PropTypes.shape({
   current: PropTypes.instanceOf(Element)
 }).isRequired,
 sendCoordinates: PropTypes.func.isRequired,
 height: PropTypes.string
};
